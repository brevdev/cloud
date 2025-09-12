package kubernetes

import (
	"context"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"

	certificatesv1 "k8s.io/api/certificates/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// ClientCertificateData creates a client certificate for the given cluster and private key. This is used to to authenticate to the cluster.
func ClientCertificateData(ctx context.Context, k8sClient *kubernetes.Clientset, username string, userPrivateKey any) ([]byte, error) {
	// Check to see if the CSR already exists for this username
	csr, err := k8sClient.CertificatesV1().CertificateSigningRequests().Get(ctx, username, metav1.GetOptions{})
	if err != nil {
		// If the error is not a not found error, return the error. If it is a not found error, continue.
		if !k8serrors.IsNotFound(err) {
			return nil, fmt.Errorf("failed to get CSR: %w", err)
		}
	} else {
		// If there is no error and the CSR exists, return the certificate.
		if csr != nil && csr.Status.Certificate != nil {
			return csr.Status.Certificate, nil
		}
	}

	// Create the certificate request
	certRequestBytes, err := x509.CreateCertificateRequest(rand.Reader, &x509.CertificateRequest{
		Subject: pkix.Name{
			CommonName:   username,
			Organization: []string{"brev"},
		},
	}, userPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create CSR: %w", err)
	}

	// Encode CSR & key in PEM
	crtRequestPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE REQUEST", Bytes: certRequestBytes})

	// Create the CSR
	csr, err = k8sClient.CertificatesV1().CertificateSigningRequests().Create(ctx,
		&certificatesv1.CertificateSigningRequest{
			ObjectMeta: metav1.ObjectMeta{Name: username},
			Spec: certificatesv1.CertificateSigningRequestSpec{
				Usages:     []certificatesv1.KeyUsage{certificatesv1.UsageClientAuth},
				Request:    crtRequestPEM,
				SignerName: certificatesv1.KubeAPIServerClientSignerName,
			},
		},
		metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to create certificate signing request: %w", err)
	}

	// Approve the CSR
	csr.Status.Conditions = append(csr.Status.Conditions, certificatesv1.CertificateSigningRequestCondition{
		Type:           certificatesv1.CertificateApproved,
		Status:         "True",
		Reason:         "BrevCloudSDK",
		Message:        "BrevCloudSDK approved certificate signing request",
		LastUpdateTime: metav1.Now(),
	})
	csr, err = k8sClient.CertificatesV1().CertificateSigningRequests().UpdateApproval(ctx, csr.Name, csr, metav1.UpdateOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to approve certificate signing request: %w", err)
	}

	// Get the signed certificate
	signedCertificate, err := k8sClient.CertificatesV1().CertificateSigningRequests().Get(ctx, csr.Name, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get signed certificate: %w", err)
	}

	return signedCertificate.Status.Certificate, nil
}

// SetUserRole sets the role for the given user.
func SetUserRole(ctx context.Context, k8sClient *kubernetes.Clientset, username string, roleName string) error {
	clusterRoleBindingName := fmt.Sprintf("%s-%s", username, roleName)

	// Check to see if the cluster role binding already exists for this username
	clusterRoleBinding, err := k8sClient.RbacV1().ClusterRoleBindings().Get(ctx, clusterRoleBindingName, metav1.GetOptions{})
	if err != nil {
		if !k8serrors.IsNotFound(err) {
			return fmt.Errorf("failed to get cluster role binding: %w", err)
		}
	} else {
		if clusterRoleBinding != nil {
			return nil
		}
	}

	// Create the cluster role binding
	_, err = k8sClient.RbacV1().ClusterRoleBindings().Create(ctx, &rbacv1.ClusterRoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name: clusterRoleBindingName,
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     roleName,
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "User",
				Name:      username,
				Namespace: "default",
			},
		},
	}, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to set role %s for user %s: %w", roleName, username, err)
	}

	return nil
}

package v1

import (
	"context"
	"net/url"
	"slices"
	"strings"
	"time"

	cloud "github.com/brevdev/cloud/v1"
)

type imageProductListResponse struct {
	Response productList `json:"getServerImageProductListResponse"`
}

func (r *imageProductListResponse) apiError() error {
	return r.Response.apiError()
}

func (c *NaverClient) GetImages(ctx context.Context, args cloud.GetImageArgs) ([]cloud.Image, error) {
	params := url.Values{}
	params.Set("regionCode", c.location)
	if len(args.Architectures) > 0 {
		for i, arch := range args.Architectures {
			if arch == architectureX8664 {
				params.Set(indexedParam("platformTypeCodeList", i+1), "LNX64")
			}
		}
	}

	var resp imageProductListResponse
	if err := c.do(ctx, "getServerImageProductList", params, &resp); err != nil {
		return nil, err
	}

	images := make([]cloud.Image, 0, len(resp.Response.ProductList))
	for _, product := range resp.Response.ProductList {
		image := cloud.Image{
			ID:           product.ProductCode,
			Name:         product.ProductName,
			Description:  firstNonEmpty(product.ProductDescription, product.OSInformation),
			Architecture: architectureFromPlatform(product.PlatformType.Code),
			CreatedAt:    time.Time{},
		}
		if len(args.ImageIDs) > 0 && !slices.Contains(args.ImageIDs, image.ID) {
			continue
		}
		if len(args.NameFilters) > 0 && !matchesAnyNameFilter(image.Name, args.NameFilters) {
			continue
		}
		if len(args.Architectures) > 0 && !slices.Contains(args.Architectures, image.Architecture) {
			continue
		}
		images = append(images, image)
	}
	return images, nil
}

func architectureFromPlatform(platform string) string {
	if strings.Contains(platform, "64") {
		return architectureX8664
	}
	if strings.Contains(platform, "32") {
		return "i386"
	}
	return ""
}

func matchesAnyNameFilter(name string, filters []string) bool {
	for _, filter := range filters {
		if strings.Contains(strings.ToLower(name), strings.ToLower(strings.Trim(filter, "*"))) {
			return true
		}
	}
	return false
}

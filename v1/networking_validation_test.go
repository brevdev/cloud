package v1

import "testing"

func TestDockerEgressPingReceivedReply(t *testing.T) {
	tests := []struct {
		name string
		out  string
		want bool
	}{
		{
			name: "all packets received",
			out: `PING 8.8.8.8 (8.8.8.8): 56 data bytes
64 bytes from 8.8.8.8: seq=0 ttl=117 time=41.193 ms
64 bytes from 8.8.8.8: seq=1 ttl=117 time=41.023 ms
64 bytes from 8.8.8.8: seq=2 ttl=117 time=41.023 ms

--- 8.8.8.8 ping statistics ---
3 packets transmitted, 3 packets received, 0% packet loss`,
			want: true,
		},
		{
			name: "partial packet loss still proves egress",
			out: `PING 8.8.8.8 (8.8.8.8): 56 data bytes
64 bytes from 8.8.8.8: seq=0 ttl=117 time=41.193 ms
64 bytes from 8.8.8.8: seq=2 ttl=117 time=41.023 ms

--- 8.8.8.8 ping statistics ---
3 packets transmitted, 2 packets received, 33% packet loss`,
			want: true,
		},
		{
			name: "iputils ping format",
			out: `--- 8.8.8.8 ping statistics ---
3 packets transmitted, 1 received, 66% packet loss, time 2003ms`,
			want: true,
		},
		{
			name: "no replies",
			out: `--- 8.8.8.8 ping statistics ---
3 packets transmitted, 0 packets received, 100% packet loss`,
			want: false,
		},
		{
			name: "no parseable stats",
			out:  `ping failed`,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dockerEgressPingReceivedReply(tt.out); got != tt.want {
				t.Fatalf("dockerEgressPingReceivedReply() = %v, want %v", got, tt.want)
			}
		})
	}
}

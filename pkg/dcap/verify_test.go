package dcap

import (
	"testing"

	"github.com/carv-protocol/verifier/internal/conf"
)

func TestVerifyAttestation(t *testing.T) {
	//t.Parallel()
	type args struct {
		data string
		cf   *conf.Bootstrap
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				data: "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",
				cf: &conf.Bootstrap{
					Dacp: &conf.Dacp{
						TrustedPath:  "../../configs/trusted.json",
						TcbPath:      "../../configs/tcb.json",
						IdentityPath: "../../configs/identity.json",
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := VerifyAttestation(tt.args.data, tt.args.cf)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyAttestation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("VerifyAttestation() = %v, want %v", got, tt.want)
			}
		})
	}
}

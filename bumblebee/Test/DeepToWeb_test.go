package Test

import (
	"bumblebee/Controllers"
	"testing"
)

func TestDeepToWeb(t *testing.T) {
	type args struct {
		deepUrl string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Deep To Web Search Test 1",
			args: args{
				deepUrl: "ty://?Page=Search&Query=elbise",
			},
			want:    "https://www.trendyol.com/sr?q=elbise",
			wantErr: false,
		},
		{
			name: "Deep To Web Search Test 2",
			args: args{
				deepUrl: "ty://?Page=Search&Query=%C3%BCt%C3%BC",
			},
			want:    "https://www.trendyol.com/sr?q=%C3%BCt%C3%BC",
			wantErr: false,
		},
		{
			name: "Deep To Web Product Test 1",
			args: args{
				deepUrl: "ty://?Page=Product&ContentId=1925865&CampaignId=439892&MerchantId=105064",
			},
			want:    "https://www.trendyol.com/brand/name-p-1925865?boutiqueId=439892&merchantId=105064",
			wantErr: false,
		},
		{
			name: "Deep To Web Product Test 2",
			args: args{
				deepUrl: "ty://?Page=Product&ContentId=1925865",
			},
			want:    "https://www.trendyol.com/brand/name-p-1925865",
			wantErr: false,
		},
		{
			name: "Deep To Web Product Test 2",
			args: args{
				deepUrl: "ty://?Page=Product&ContentId=1925865&MerchantId=105064",
			},
			want:    "https://www.trendyol.com/brand/name-p-1925865?merchantId=105064",
			wantErr: false,
		},
		{
			name: "Deep To Web Product Test 3",
			args: args{
				deepUrl: "ty://?Page=Product&ContentId=1925865&CampaignId=439892",
			},
			want:    "https://www.trendyol.com/brand/name-p-1925865?boutiqueId=439892",
			wantErr: false,
		},
		{
			name: "Deep To Web Home Page Test 1",
			args: args{
				deepUrl: "ty://?Page=Favorites",
			},
			want:    "https://www.trendyol.com",
			wantErr: false,
		},
		{
			name: "Deep To Web Home Page Test 1",
			args: args{
				deepUrl: "ty://?Page=Orders",
			},
			want:    "https://www.trendyol.com",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Controllers.DeepToWeb(tt.args.deepUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeepToWeb() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DeepToWeb() = %v, want %v", got, tt.want)
			}
		})
	}
}

package Test

import (
	"bumblebee/Controllers"
	"testing"
)

func TestWebToDeep(t *testing.T) {
	type args struct {
		webUrl string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Search Url test case 1",
			args: args{
				webUrl: "https://www.trendyol.com/sr?q=elbise",
			},
			want:    "ty://?Page=Search&Query=elbise",
			wantErr: false,
		},
		{
			name: "Search Url test case 2",
			args: args{
				webUrl: "https://www.trendyol.com/sr?q=%C3%BCt%C3%BC",
			},
			want:    "ty://?Page=Search&Query=%C3%BCt%C3%BC",
			wantErr: false,
		},
		{
			name: "Product Url test case 1",
			args: args{
				webUrl: "https://www.trendyol.com/casio/saat-p-1925865?boutiqueId=439892&merchantId=105064",
			},
			want:    "ty://?Page=Product&ContentId=1925865&CampaignId=439892&MerchantId=105064",
			wantErr: false,
		},
		{
			name: "Product Url test case 2",
			args: args{
				webUrl: "https://www.trendyol.com/casio/erkek-kol-saati-p-1925865",
			},
			want:    "ty://?Page=Product&ContentId=1925865",
			wantErr: false,
		},
		{
			name: "Product Url test case 3",
			args: args{
				webUrl: "https://www.trendyol.com/casio/erkek-kol-saati-p-1925865?boutiqueId=439892",
			},
			want:    "ty://?Page=Product&ContentId=1925865&CampaignId=439892",
			wantErr: false,
		},
		{
			name: "Product Url test case 4",
			args: args{
				webUrl: "https://www.trendyol.com/casio/erkek-kol-saati-p-1925865?merchantId=105064",
			},
			want:    "ty://?Page=Product&ContentId=1925865&MerchantId=105064",
			wantErr: false,
		},
		{
			name: "Home Page Test Case 1",
			args: args{
				webUrl: "https://www.trendyol.com/Hesabim/Favoriler",
			},
			want:    "ty://?Page=Home",
			wantErr: false,
		},
		{
			name: "Home Page Test Case 2",
			args: args{
				webUrl: "https://www.trendyol.com/Hesabim/#/Siparislerim",
			},
			want:    "ty://?Page=Home",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Controllers.WebToDeep(tt.args.webUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("WebToDeep() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("WebToDeep() = %v, want %v", got, tt.want)
			}
		})
	}
}

package cloudinary

import "github.com/cloudinary/cloudinary-go/v2"

func InitCloudinary() *cloudinary.Cloudinary {
	cld, err := cloudinary.NewFromParams("dckghjaen", "284443345373777", "kone0OvtBr_rPZr_e-6n9rQe9pc")
	if err != nil {
		panic(err)
	}
	return cld
}

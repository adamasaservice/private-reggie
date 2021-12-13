package main

// List Available Versions
// https://www.terraform.io/docs/internals/provider-registry-protocol.html
type VersionResp struct {
	Version   string   `json:"version"`
	Protocols []string `json:"protocols"`
	Platforms []struct {
		Os   string `json:"os"`
		Arch string `json:"arch"`
	} `json:"platforms"`
}
type Versions struct {
	Versions []VersionResp `json:"versions"`
}

// Find a Provider Package
// https://www.terraform.io/docs/internals/provider-registry-protocol.html
type DownloadResp struct {
	Protocols           []string `json:"protocols"`
	Os                  string   `json:"os"`
	Arch                string   `json:"arch"`
	Filename            string   `json:"filename"`
	DownloadURL         string   `json:"download_url"`
	ShasumsURL          string   `json:"shasums_url"`
	ShasumsSignatureURL string   `json:"shasums_signature_url"`
	Shasum              string   `json:"shasum"`
	SigningKeys         struct {
		GpgPublicKeys []struct {
			KeyID          string `json:"key_id"`
			ASCIIArmor     string `json:"ascii_armor"`
			TrustSignature string `json:"trust_signature"`
			Source         string `json:"source"`
			SourceURL      string `json:"source_url"`
		} `json:"gpg_public_keys"`
	} `json:"signing_keys"`
}

// config to load
type Config []struct {
	Name      string   `json:"name"`
	Protocols []string `json:"protocols"`
	Versions  []string `json:"versions"`
	Platforms []struct {
		Os   string `json:"os"`
		Arch string `json:"arch"`
	} `json:"platforms"`
	Os                  string `json:"os"`
	Arch                string `json:"arch"`
	Filename            string `json:"filename"`
	DownloadURL         string `json:"download_url"`
	ShasumsURL          string `json:"shasums_url"`
	ShasumsSignatureURL string `json:"shasums_signature_url"`
	SigningKeys         struct {
		GpgPublicKeys []struct {
			KeyID          string `json:"key_id"`
			ASCIIArmor     string `json:"ascii_armor"`
			TrustSignature string `json:"trust_signature"`
			Source         string `json:"source"`
			SourceURL      string `json:"source_url"`
		} `json:"gpg_public_keys"`
	} `json:"signing_keys"`
}

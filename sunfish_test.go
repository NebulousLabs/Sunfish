package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var (
	logDir = "testLogs"
	dbName = "testSunfish"
	sf     *Sunfish
	url    string
)

// TestGetSiafiles test whether the api endpoint returns an empty list of
// siafiles to start
func TestGetSiafiles(t *testing.T) {
	var siafiles []Siafile
	req, err := http.NewRequest("GET", url+"/siafile/", nil)

	if err != nil {
		t.Errorf("Error: problem creating request: %s", err)
		sf.logger.Println(err)
	}

	res, err := http.DefaultClient.Do(req)

	if res.StatusCode != 200 {
		t.Errorf("Error: expected 200. Received: %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(body, &siafiles)
	if err != nil {
		t.Errorf("Error unmarshalling JSON into siafiles. %s", err)
	}

	if len(siafiles) != 0 {
		t.Errorf("Expected 0 Siafiles. Found %d", len(siafiles))
	}
}

func TestAddSiafile(t *testing.T) {
	var siafile Siafile
	siafileString := `{
		"listed":true,
		"safe":true,
		"title":"Title",
		"description":"description",
		"ascii":"H4sIAAAJbogC_-xdW3PcNpZ-ln9FVz9n27hf_JbKpZLajeOyk92HlGsLBMGox61uT1-S8Uzlvw9bEs933KCUi6UoGlMPDgKCAAiA-L7znQP2_NUyzV6dp21pZ18uV2X-ZC4Wcv7khydn_3pydjZ_ni7K_Nlsnjeb1f_vNusfF90q5fknx2ufnZf8Zne46K_3xc_OxCeP9t_-n9eXz_TFNu0O2_Iqn5erB7961BfLksvuZfn7YdmPVJ8vL7O_fbtfXqTVy5I3P5Xtu6tS_dXLWuffbfZpdZJH_3s5YschPhbcpvWuK9urqmNQxpgYPrm6-GneL386dmW_PZTrvJflbVpul-sf--wurXZD_meb9X6b8r7Pvqr6bH6c1FfLf5aq4qtL35Ttm1V5udnshz71f0qr6zJnUtkhqaQfkpryqCDKSSWGZDSUGaikoZIOzQhHmQYNClwfen3mKdOgpKFMpSnp5EjXpUVR49A9qj-iU4YqMPTEAY2izwrjICOaUvIq9XoY8f9brtvNz6_2aXscbqkDNXZ96Yv15eIyMgy9nL9I7zaHY3EbonEuhiAEXfzftFq2L7abTfftYf_2sN-xWbxeAVelDscFYJ0zvv9jNfRXv1-vNvnNV2l3flzvKVudtQnZNso2nU2tLanL0RubW-NVX4HoorNddl7l5GOrddfkxumkWqdyMTam3KlsfZhft_HLyTB8s9ztyt12W3zgX4ilEa1LIbl4U7fvscWhiZflp-VuuVk_P1w0ZXvcNS4v_HLyfn_9OQYMr5uiNwsrWw9LWPphhbthgUo9rGoph4sqXifskJCCakV5eh2VppRRlDe8I0pQ81SxQpeUo3upT4HajUNrQdKtiro-NEEvrnV0o6BuUmV0p5L2CZvV-Veb3f7rF8cJVSIu3MKFhRL6Wew3y-tZmV--r1-v2_KPYRs_AsW6Pc263NtH99qrS3UVefuuR5DN-r_LO0yoGsaCdhJlNQ3oMImBnnYoRTOBeVUaA2DpBkNLgsaf9kdsn1gStLVjFKXTVWM0E1IMLaABFapnoYtS6apN6iN1h7Z46Sj13kye0pEB3D_C_16NyxPaOf6yPEMCgAHLFvQCCAvYtuAPVFDjFsnwG_gM0gBSQGuyJwVxhIlIwXgOkqGuyqKfmnEKDVIAIuQiCARyPT2UQa8ihkegg86O0R6jJq4xcY375hrR1YzBAB6G9epDBfZgHRI8gVCNoMPSxk9ApPQIDhJISjQvAHGyrhmIKYmA0Muqa4YRAE_Dnea0OFgFdZveTinNDUTjh2fP5OuHZBhOnM4AyFZvtNDYqVOktjQTZBCBHFCtztZkwp-OmKRtnLYfKWnIh4Q3IxMeq2mmKhToQqhSNNtYxfQ84E7UoJvYxaNmF1KoMfsc-7r0gGJGIADwLFd6BrUMzIkBUF4AASCrhosagvETdEWyvgLqwX8sK8pkDagmSI30Ex3xuCw86om1ZiKZ5uHui1tc_30YxbDG9XVJcQtkt6Lpi7TWtyG7futpRKdaLWVsc1Ncl9vSlqicskXG0LiSmzb76JU0XdeU6HTOojT9Fdsv7rtlGr-h9x8v4SBTnqFSZbNKIvrQAMAKgA0gAMNyjwQlAHALOIJZSrWZCrSGBLPRiRVQr0nCkMZViKbkUH20Ne5RJ_wIqRGQRsINbMPZRXQL6eQi2p50mIciHQTDlGAahq5YH4ggTQ2xDzZbUBggNTHOpypWI1EL7Yo0vVZWmgcMTwgXqiYsSlVsEUYnxBdVMyOwVpK-lJ-ox-OmHjThZ54pE2MiBVwtEFC5M4VZ-xrEgFEQJUcaAIrTDgp9kLERhvFoR4wJIJBqtBmRTbDVnaGXtK-dYTwYfZHwynB9hT3F3TpQenPPWOFivxcGp53ojb0_zD16W1eq2CNPX6UVsWdQ6jb64bpS-qlstHWyaaVNwTQl25yEl74tOsiehOjkbAy-CV1RTeqUNjnqrHq4akxPmWIbWyfuin78vgf4eBnIsNhJhABBIPagAFYki8AXAyiht5ZeQCga2P2BRrbS-OHsgX-GQItojZKVEAMxEl5NX2cB7eBDAA0iGwgeFhL94038o6cfCykXITyk4gEGB68t1B_G-WiyQE5IZVan8g82GQnnC80oqod2ESE9UR5TxkbIDNEEUTXAlTeQCFXNuK-fTqE65nobVuW1ljzxj0fLPxSCEMAqmIbAvCDM2wLgRYwGaWU8mMMzzlETiABYZ1wgjvhFuEThxhgN2BHnQWHMgwJ2ZJk-o0c6xXoFxw3z1bCK_jPohxdelNB45bpWdC6pKHxrg1WmiznYkrNzQXadDE1jihQp2H6riD4lWZrWZi9ULDEEJ9JEP_5c-qGYG0RUJABhEAAgSwlRRVcowgn40OENofdGkYkObgFXhxC1FKNGRA1XBRxEVwkj8O9LU3cX9ji8_96eeh5AcrS8gYRoueiX20I6t9APy0OIQ2DmYuU7IiRGKa_qUAppR-ZkaIC2eOlcFUNBMi8pL0QGSR9BiE2s-ANNqYtVEAebeVLCiMpijUmoPyOEd2Igj5uBsDAIwDWPiADyWhYxwWiLG5EWIOi9F30JaGeCBdCEaSWMj4iRIBOmQrAwDMRmgCsxrUYwsQVRpNQS8iwKQhVCM4rRFjyRvh8K4qyPWgfVb6wfzkGiCMobI4y5BcObEFoVcnElm9Z0tpimU21JKbrsfb8eWhe0i23wOfUw4mKKtnWq5J6V-BCNVcWbPtv4ePck5Dc9wUcsgohTE9iYES-JPoUeFrChYyUxAAzx4iF01IwEm9pad4BtzYgBnDkkx-hTesRsbISrELa6ETcT8z_Umj2CHvte3iKFHFnIw8aXMjkCw4Nhl5XgBZkBCtOJPAEPMhMqYh0tgrBRIo9gdjSClEVyCdgnU9PUqVMOLmtYqOAg4KGeHD8QZijihFjlFF36yEUQSBfsJAkLJEWcOvdP4B74cOyISoAwDhwJAW4jtAMOGB6FwkgPC1gNckS54GIOI1A8YNaMBJpCHeTaCvH9My9H4mSZngLnjPW_ykH0TRxE1hzEeO1cD-XvQ9Uf5CBGuv6ZjQ7uur5bAymCUzJn2XrfGmOz0a2TRaUgfPZaZ1tc6nxJvt-wTGliSNq4rFRxoet7bJORthWNd6ULdxVw-juf4CNWQmjnB5bXPndRGbZMWFAVjQHxAJqQYcGgBrB1Kq7AkKCX39WxoCNRkvAdQJ8nYFIjTgoYzLYyq2E89HbRDezDm0UQi35bWMQH9cXQzsnIYR3QEmtHDJ4WQ4EITlULVEy0wppxtYpGPhOoGRSAw1w24BQRzr1T1QSyizgN-oDVZ0bccgQyOFAzEZBHTUBMHPG3sEBSfr7Ds0MdeuyAamSHVdgJETcaXSJHYkYY24m2dpUYFuo6dgKWncDl_ab7_RiXCnHk2C2XdgxjRHLkgDCq-lNIyJ2xkNtBXDdadTlpEYNIbQ4hSRVjJ0PRSZpctNI-yGyzkq647FvpWt1GrdssU8qqNCk740Ir_f3RkImH3MBD6CiWC5XR7OJIYIeppAccYYBTh9a6cZUugaOSBFW6FvXx3uO9NSMBh77WSZypYAjnJsBP2KEfhCNUjiRYWyeBBDcGpboHOwlTH4UmXsbEBF-JXSwqhjRnGmLE_ohTegD64eszK5AjhK9Dc6Q6HXKiTZhcXR-B0qc-HzbHiB8mkZ02cra8p3jUD2EiQQVjiPiDiXDKwakIYyi3MxHxqxSE9u8p8V7iVi4hxnhE7Uy5lSysD6vVrVDMC_w1ge4jf6krzHogiJqmYcKcxxmDGOpPJvBvLrAYRJiRsFfxHSXFrFCEG6oR25kF-Um43X2l7TML2o2p6JD-ccCIxz8ijmA0DgAmAvswBRf32RlR9jUI1qkAY_pXzV5zk9mrps86TJ91-E2nLE0louLckK7sUkicZFsgso9MWxVrI2MkHICJ5GQ8I0iMNCtmBsnKG43Td_DsmkqrZSF0JM3iLIQ6DXaj_gdY4X_VDzqYKnwCDhFfH4vEl5qYpRoqNzlsXMRVQZOAByTikwmVh4KvKHwBDD4NW6kk2CdV5ZxHcbKeKSIMX6TAiUpbLwI1RRbeB6-YbNnJlp1s2cmWnWzZyZa9a8xRVx9JmTBnwpwJcybMmTBnwpx7x5zJWTeBzQQ2E9hMYDOBzQQ2E9hMYDOBzQQ2E9hMYDOBzQQ2E9hMYDOBzTQu12BDv575_dvVJrUv0jZd7AasuESF9fXPiD69KO0yPb047Jb56TfpfFW2s1fvLt6eb9bLspt9-T-ffvZUzt4cmrJavnkqhJk93yxmcrZczz6fXaS_bbaz75b7tJ7918wsZq_2h-3FcpfPZ035ufy4x2-Tns0_P2zTcRH3rboh6Gv-fJnfrG_-RVP285yXvxLXP9gvT17_OwAA__8O4tKkI3UAAA==",
		"filename":"m1-4.sia",
		"tags":["tag","test"]}`
	req, err := http.NewRequest("POST", url+"siafile/", strings.NewReader(siafileString))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Errorf("%s", err)
	}

	res, err := http.DefaultClient.Do(req)

	if res.StatusCode != 201 {
		t.Errorf("Error: expected 201. Received: %d, %s", res.StatusCode, url+"siafile/")
	}

	if err != nil {
		t.Errorf("%s", err)
	}

	body, err := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(body, &siafile)
	if err != nil {
		t.Errorf("Error unmarshalling JSON into siafile. Error: %s", err)
	}

	if siafile.Title != "Title" {
		t.Errorf("Title Error: Expected 'Title' got %s", siafile.Title)
	}

	if siafile.Description != "description" {
		t.Errorf("Description Error: Expected 'description' got %s", siafile.Description)
	}
}

func TestMain(m *testing.M) {
	sf = NewSunfish(logDir, dbName)

	// Drop db so we have a clean db for testing
	sf.DB.DropDatabase()
	server := httptest.NewServer(sf.Router)

	url = server.URL + "/api/"

	err := m.Run()

	sf.Close()
	server.Close()
	os.Exit(err)
}

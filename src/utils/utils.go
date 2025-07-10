package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"mime/multipart"
	"net/http"
	"reflect"
	"strings"
	"time"
	"unicode"
)

func BeautifyJson(v interface{}) string {
	jsonBytes, err := json.MarshalIndent(v, "", "  ")

	if err != nil {
		return ""
	}

	return string(jsonBytes)
}

func MergeDefaults[T any](target T, defaults T) T {
	targetVal := reflect.ValueOf(&target).Elem()
	defaultVal := reflect.ValueOf(defaults)

	for i := range targetVal.NumField() {
		field := targetVal.Field(i)
		defaultField := defaultVal.Field(i)

		if IsZero(field) && field.CanSet() {
			field.Set(defaultField)
		}
	}

	return target
}

func IsZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Slice, reflect.Map, reflect.Ptr, reflect.Interface:
		return v.IsNil()
	case reflect.Struct:
		return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
	default:
		return false
	}
}

func UploadImageBytesToUploaderSh(data []byte) (string, error) {
	fileName := fmt.Sprintf("screenshot-%s-%d.png", time.Now().Format("20060102-150405"), rand.Intn(math.MaxInt32))

	return UploadBytesToUploaderSh(data, fileName)
}

func UploadBytesToUploaderSh(data []byte, fileName string) (string, error) {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return "", fmt.Errorf("error creando form file: %w", err)
	}

	if _, err := part.Write(data); err != nil {
		return "", fmt.Errorf("error escribiendo datos: %w", err)
	}

	writer.Close()

	req, err := http.NewRequest("POST", "https://uploader.sh/", &body)
	if err != nil {
		return "", fmt.Errorf("error creando request: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error en la request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("uploader.sh devolvió estado %d", resp.StatusCode)
	}

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error leyendo respuesta: %w", err)
	}

	return GetResponseUrlFromUploaderSh(string(result), "wget "), nil
}

func GetResponseUrlFromUploaderSh(text string, toSearch string) string {
	position := strings.Index(text, toSearch)

	if position == -1 {
		return text
	}

	return text[position+len(toSearch):]
}

func SlugifyUpper(input string) string {
	replacer := strings.NewReplacer(
		"á", "a", "é", "e", "í", "i", "ó", "o", "ú", "u",
		"Á", "A", "É", "E", "Í", "I", "Ó", "O", "Ú", "U",
		"ñ", "n", "Ñ", "N",
	)

	normalized := replacer.Replace(input)

	normalized = strings.ReplaceAll(normalized, " ", "_")

	cleaned := strings.Builder{}
	for _, r := range normalized {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
			cleaned.WriteRune(r)
		}
	}

	return strings.ToUpper(cleaned.String())
}

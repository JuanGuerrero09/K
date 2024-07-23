package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

func encryptFile(key, inputFile, outputFile string) error {
	// Read the plaintext file
	plaintext, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	// Crear un nuevo cifrado AES usando la llave
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}

	// Crear un nuevo GCM (Galois/Counter Mode) cifrado
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// Generar un nuevo IV (vector de inicialización)
	iv := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	// Encrypt the plaintext
	ciphertext := gcm.Seal(iv, iv, plaintext, nil)

	// Write the ciphertext to the output file
	return os.WriteFile(outputFile, ciphertext, 0644)
}

func decryptFile(key, inputFile, outputFile string) error {
	// Leer el archivo cifrado
	ciphertext, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	// Crear un nuevo cifrado AES usando la llave
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return err
	}

	// Crear un nuevo GCM (Galois/Counter Mode) cifrado
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// Extraer el IV del comienzo del texto cifrado
	iv := ciphertext[:gcm.NonceSize()]
	ciphertext = ciphertext[gcm.NonceSize():]

	// Descifrar el texto cifrado
	plaintext, err := gcm.Open(nil, iv, ciphertext, nil)
	if err != nil {
		return err
	}

	// Escribir el texto en claro en el archivo de salida
	return os.WriteFile(outputFile, plaintext, 0644)
}

	
func check(e error) {
	if e != nil {
			panic(e)
	}
}

func getEncryptedTodos(key string) (todos []ToDo, temp *os.File){
	err := encryptFile(key, "list.txt", "newfile2.txt")
	check(err)
	f, err := os.CreateTemp("", "sample")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil
	}
	defer os.Remove(f.Name())
	err = decryptFile(key, "newfile2.txt", f.Name())
	check(err)
	file, err := os.Open(f.Name())
	if err != nil {
		fmt.Println("Error creating temp file:", err)
		return nil, nil
	}
	todos, err = ParseFile(file)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return nil, nil
	}
	return todos, f
}

func encryptTodos(key string, todos []ToDo, f *os.File) {
	saveFile(f.Name(), todos)
	err := encryptFile(key, f.Name(), "newfile2.txt")
	check(err)
	err = decryptFile(key, "newfile2.txt", "list.txt")
	check(err)
}
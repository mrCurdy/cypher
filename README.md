# Cypher

## Description

<!-- Add third encription -->

This command-line tool allows users to encrypt and decrypt messages using various encryption techniques. The tool supports two types of encryption: ROT13 and Reverse Alphabet. It is designed to handle user inputs, ensuring that only valid inputs are processed. All non-alphabet characters are preserved in the message during encryption and decryption operations.

## Features

Greet the user: The tool welcomes the user upon start.
Select Operation: The user can choose to either encrypt or decrypt a message.
Select Encryption Type: The user can choose between ROT13 and Reverse Alphabet types of encryption.
Input Message: The user inputs the message that is going to be encrypted or decrypted.
Output Result: The tool outputs the encrypted or decrypted message.
Input Validation: The tool continuously prompts the user for valid input until it receives it. Leading and trailing whitespaces are removed before validation.

## Encryption Techniques

ROT13: This technique shifts each letter of the alphabet by 13 places. The process is the same for encryption and decryption.

Example:

Input: Hello World!
Output: Uryyb Jbeyq!

Reverse Alphabet: This technique reverses the position of each letter in the alphabet.

Example:

Input: Hello World!
Output: Svool Dliow!

## Usage

Run the code:
Follow the on-screen promts:
First you will be greeted by the tool. 

<!-- add insert number -->

Select whether you want to encrypt or decrypt the message.
Select between ROT13 or Reverse Alphabet encryption/decryption techniques.



Choose the encryption technique: ROT13 or Reverse Alphabet.
Input the message you want to encrypt or decrypt.
The tool will display the result of the operation.
Example session:

Copiar código
Welcome to the Encryption CLI Tool!
Please select an operation (encrypt/decrypt): encrypt
Please select the encryption type (rot13/reverse): rot13
Please enter the message: Hello World!
Encrypted message: Uryyb Jbeyq!

Encryption Techniques
What does the tool do?
Tool usage with examples.
Explanation of the cyphers used. ...
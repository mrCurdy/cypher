# Cypher

## Description

<!-- Add third encription -->

This command-line tool allows users to encrypt and decrypt messages using various encryption techniques. The tool supports three types of encryption: ROT13, Reverse and Reverse and shift. It is designed to handle user inputs, ensuring that only valid inputs are processed. All non-alphabet characters are preserved in the message during encryption and decryption operations.

## Features

Greet the user: The tool welcomes the user upon start.
Select Operation: The user can choose to either encrypt or decrypt a message.
Select Encryption Type: The user can choose between ROT13, Reverse and Reverse and shift types of encryption.
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

Reverse and shift: This technique reverses the string of characters and then shifts it by 13 places.

Example:

Input: Hello World!
Output: !qyebJ byyrU

## Usage

Run the code:
Follow the on-screen promts:
First you will be greeted by the tool. 

Select whether you want to encrypt or decrypt the message, by entering the number 1 or 2 accordingly.

Select between ROT13, Reverse and Reverse and shift encryption/decryption techniques, by entering the number 1, 2 or 3 accordingly.

Enter the message you want to encrypt or decrypt.

The tool will display the result of the operation.

## Example
Example session:

Welcome to the Encryption Cypher Tool!

Select operation (1/2):
1. Encrypt.
2. Decrypt.
2

Select cypher (1/3):
1. ROT13
2. Reverse
3. Reverse and shift
3

Please enter the message: 
viubW/qbbx

Decrypted message using Reverse and shift: 
kood/Johvi

Thank you for your time!
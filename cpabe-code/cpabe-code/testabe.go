/*
 * Copyright (c) 2018 XLAB d.o.o
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */


package main

import (
    "fmt"
    "./cpabe"


)

func main() {


	policy := "((0 AND 1) OR (2 AND 3)) AND 5"
	fmt.Println("policy =>")
	fmt.Println(policy)

	// define a set of attributes (a subset of the universe of attributes)
	// that an entity possesses
	gamma := []int{0, 2, 3,5}
	fmt.Println("gamma =>")
	fmt.Println(gamma)	

	// create a message to be encrypted
	msg := "Attack at dawn!"
	fmt.Println("msg =>")
	fmt.Println(msg)

	// create a new FAME struct with the universe of attributes
	// denoted by integer
	a := abe.NewFAME()

	// generate a public key and a secret key for the scheme
	pubKey, secKey, _ := a.GenerateMasterKeys()

	// create a msp struct out of a boolean expression representing the
	// policy specifying which attributes are needed to decrypt the ciphertext;
	// note that safety of the encryption is only proved if the mapping
	// msp.RowToAttrib from the rows of msp.Mat to attributes is injective, i.e.
	// only boolean expressions in which each attribute appears at most once
	// are allowed - if expressions with multiple appearances of an attribute
	// are needed, then this attribute can be split into more sub-attributes

	msp, _:= abe.BooleanToMSP(policy, false)

	// encrypt the message msg with the decryption policy specified by the
	// msp structure

	cipher, _ := a.Encrypt(msg, msp, pubKey)

	fmt.Println("msg =>")
	fmt.Println(msg)
	fmt.Println("Encryption...")
 	fmt.Printf("Ciphertext: %+v\n ", cipher)


	// generate keys for decryption for an entity with
	// attributes gamma
	keys, _:= a.GenerateAttribKeys(gamma, secKey)
	fmt.Println(keys)	



	//*******************************************************//
	// decrypt the ciphertext with the keys of an entity
	// that has sufficient attributes

	token, _ := a.PreDecrypt(cipher, keys, pubKey) 

	msgCheck, _ := a.Decrypt(cipher, token,secKey.Zu)

	fmt.Println("Decrypt...")
	fmt.Println(msgCheck)
    if msg == msgCheck { 
    	fmt.Println("Successful Decryption!!!")
    }else{
		fmt.Println("Decryption. Failed!!!")
    	}


}

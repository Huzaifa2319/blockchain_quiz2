//Huzaifa Anjum 20i-2319 CS-8A 
// Quiz 2
//Mam I have writen this code with the help of internet(honestly speaking) but i have a good hands on experience on git, git hub and devops
//
package main

import (
    "fmt"
)

type Block struct {
    Data string
}

type Blockchain struct {
    Blocks []*Block
}


func (bc *Blockchain) DisplayAllBlocks() {
    fmt.Println("Blocks in the blockchain:")
    for _, block := range bc.Blocks {
        fmt.Println(block.Data)
    }
}


func (bc *Blockchain) NewBlock(data string) {
    newBlock := &Block{Data: data}
    bc.Blocks = append(bc.Blocks, newBlock)
    fmt.Println("New block added to the blockchain with data:", data)
}

func (bc *Blockchain) ModifyBlock(index int, newData string) {
    if index >= 0 && index < len(bc.Blocks) {
        bc.Blocks[index].Data = newData
        fmt.Println("Block at index", index, "modified with new data:", newData)
    } else {
        fmt.Println("Invalid index")
    }
}

func main() {
    blockchain := &Blockchain{}

    blockchain.NewBlock("Block 1")
    blockchain.NewBlock("Block 2")
    blockchain.NewBlock("Block 3")
    blockchain.DisplayAllBlocks()
    blockchain.ModifyBlock(1, "Modified Block 2")
    blockchain.DisplayAllBlocks()
}

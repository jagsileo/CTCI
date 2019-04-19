package sumLists

import "github.com/CTCI/Golang/Chapter2/singleLinkedList"

func sumListsWhenNumReversed(listA *singleLinkedList.Node, listB *singleLinkedList.Node) *singleLinkedList.Node {
	a := listA
	b := listB
	result := singleLinkedList.InstantiateLinkedList()
	c := &result
	var quotient, sum int8

	for a != nil || b != nil || quotient != 0 {
		sum += quotient
		if a != nil {
			sum += a.Value
			a = a.Next
		}
		if b != nil {
			sum += b.Value
			b = b.Next
		}

		c.Next = &singleLinkedList.Node{
			Value: sum % 10,
			Next:  nil,
		}
		quotient = sum / 10
		c = c.Next
		sum = 0
	}
	return result.Next
}

func sumListsWhenNumForward(listA *singleLinkedList.Node, listB *singleLinkedList.Node) *singleLinkedList.Node {
	a := listA
	b := listB
	result := singleLinkedList.InstantiateLinkedList()
	c := &result
	var numA, numB, lenA, lenB, sum int8

	for a != nil {
		lenA++
		a = a.Next
	}

	for b != nil {
		lenB++
		b = b.Next
	}
	a = listA
	b = listB
	bufA := lenA
	bufB := lenB

	for a != nil {
		numA += a.Value * getMultipleOfTen(bufA)
		bufA--
		a = a.Next
	}

	for b != nil {
		numB += b.Value * getMultipleOfTen(bufB)
		bufB--
		b = b.Next
	}

	sum = numA + numB

	var dividend, bufC int8
	dividend = sum
	bufC = 0
	for dividend > 0 {
		dividend /= 10
		bufC++
	}

	dividend = sum

	for bufC > 0 {
		quotient := dividend / getMultipleOfTen(bufC)
		c = c.InsertElement(quotient)
		dividend -= quotient * getMultipleOfTen(bufC)
		bufC--
	}

	return c

}

func getMultipleOfTen(len int8) int8 {
	var multiple, i int8
	multiple = 1
	for i = 1; i < len; i++ {
		multiple = multiple * 10
	}
	return multiple
}

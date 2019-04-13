package myLinkedList;
import java.util.*;
public class MyLinkedList {

	public	MyNode head;

	public MyLinkedList() {
		head = null;
	}

	public MyLinkedList(int d) {
		head = new MyNode(d);
		head.next = null;
	}

	public Boolean isEmpty() {
		return head == null;
	}

	public void addFirst(int d) {
		MyNode n = new MyNode(d);
		n.next = head;
		head = n;
	}

	public void addLast(int d) {
		MyNode n = new MyNode(d);
		if(head == null) {
			head = n;
		}
		else {
			MyNode curr = head;

			while(curr.next != null) {
				curr = curr.next;
			}
			curr.next = n;
		}
	}

	public void replaceLast(int d) {
		if((head == null) || (head.next == null)) {
			head = new MyNode(d);
			head.next = null;
		}

		else {
			MyNode curr = head;
			while(curr.next != null) {
				curr =curr.next;
			}
			curr = new MyNode(d);
		}
	}
	public void printList(MyNode head) {
		MyNode curr = head;
		while(curr != null) {
			System.out.println(curr.data);
			curr = curr.next;
		}
	}
}











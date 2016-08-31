import myLinkedList.*;
import java.util.*;
public class ListKtoN {
	public void printKtoN(MyNode n) {

		for(int i = 0; i<3; i++) {
			n = n.next;
		}

		while(n != null) {
			System.out.println(n.data);
			n = n.next;
		}
	}


	public static void main(String[] args) {
		MyLinkedList ll = new MyLinkedList();
		ListKtoN kn = new ListKtoN();
		ll.addLast(5);
		ll.addLast(10);
		ll.addLast(15);
		ll.addLast(20);
		ll.addLast(25);
		ll.addLast(30);
		kn.printKtoN(ll.head);
	}
}


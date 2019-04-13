import myLinkedList.*;
import java.util.*;
public class IntersectLL {
	private Boolean checkIntersectCostly(MyNode n1, MyNode n2) {
		Boolean result = false;
		if((n2 == null)&&(n1 == null))
			return result;
		MyNode curr1 = n1;

		while(curr1 != null) {
			MyNode curr2 = n2;
			while(curr2 != null) {
				if(curr1 == curr2) {
					result = true;
					break;
				}
				curr2 = curr2.next;
			}
			curr1 = curr1.next;
		}
		return result;

	}

	private Boolean checkIntersectCheap(MyNode n1, MyNode n2) {
		MyNode curr1 = n1;
		MyNode curr2 = n2;

		while(curr1.next != null) {
			curr1 = curr1.next;
		}
		while(curr2.next != null) {
			curr2 = curr2.next;	
		}
		if(curr1 == curr2) {
			return true;
		}
		else {
			return false;
		}
	}
	public static void main(String[] args) {
		MyLinkedList ll1 = new MyLinkedList();
		MyLinkedList ll2 = new MyLinkedList();
		IntersectLL il = new IntersectLL();
		ll1.addLast(1);
		ll1.addLast(3);
		ll1.addLast(5);
		ll1.addLast(7);
		ll1.addLast(9);
		ll2.addLast(2);
		ll2.addLast(4);		
		//MyNode curr1 = ll1.head;
		//MyNode curr2 = ll2.head;
		//for(int i = 0; i<2 ; i++) {
		//	if(curr1.next != null)
		//		curr1 = curr1.next;
		//}
		//while(curr2.next != null) {
		//	curr2 = curr2.next;
		//}

		//curr2.next = curr1.next;

		ll1.printList(ll1.head);
		ll2.printList(ll2.head);
		System.out.println(il.checkIntersectCostly(ll1.head, ll2.head));
		System.out.println(il.checkIntersectCheap(ll1.head, ll2.head));
	}
}























import java.util.*;
import myLinkedList.*;
public class DetectLoopLL {
	private Boolean checkLoop(MyNode n1) {
		if(n1 == null)
			return false;

		MyNode slow = n1;
		MyNode fast = n1;		

		while(true) {

			slow = slow.next;          // 1 hop.

			if(fast.next != null)
				fast = fast.next.next; // 2 hops.
			else
				return false;          // next node null => no loop.

			if(slow == null || fast == null) // if either hits null..no loop.
				return false;

			if(slow == fast) // if the two ever meet...we must have a loop.
				return true;
		}	
	}


	public static void main(String[] args) {
		MyLinkedList ll1 = new MyLinkedList();
		DetectLoopLL dl = new DetectLoopLL();
		ll1.addLast(1);
		ll1.addLast(3);
		ll1.addLast(5);
		ll1.addLast(7);
		ll1.addLast(9);
		ll1.addLast(2);
		ll1.addLast(4);
		MyNode curr1 = ll1.head;
		for(int i = 0; i<4; i++) {
			curr1 = curr1.next;
		}
		MyNode last = ll1.head;

		while(last.next != null) {
			last = last.next;		
		}
		if(last.next == null) {
			last.next = curr1;
		}
		//		ll1.printList(ll1.head);
		System.out.println(dl.checkLoop(ll1.head));
	}
}
















import myLinkedList.*;
import java.util.*;
import java.lang.*;
public class SumIfForward {

	private MyLinkedList sumFrwd(MyNode n1, MyNode n2) {
		MyLinkedList res = new MyLinkedList();
		return res;
	}

	public static void main(String[] args) {
		MyLinkedList ll1 = new MyLinkedList();
		MyLinkedList ll2 = new MyLinkedList();
		SumIfForward s = new SumIfForward();
		ll1.addLast(9);
		ll1.addLast(1);
		ll1.addLast(5);
		ll2.addLast(5);
		ll2.addLast(7);
		ll2.addLast(3);
		ll2.addLast(8);	
		MyLinkedList result = s.sumFrwd(ll1.head, ll2.head);
		result.printList(result.head);

	}
}
















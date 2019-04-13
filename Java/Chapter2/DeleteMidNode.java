import myLinkedList.*;
import java.util.*;
public class DeleteMidNode {

	private void delMidNode(MyNode n) {
		while(n.next !=null) {
			n.data = n.next.data;
			if(n.next.next == null) {
				n.next = null;
			}
			else {
				n = n.next;
			}
		}				
	}

	public static void main(String[] args) {
		MyLinkedList ml = new MyLinkedList();
		DeleteMidNode dm = new DeleteMidNode();
		ml.addLast(5);
		ml.addLast(10);
		ml.addLast(15);
		ml.addLast(20);
		ml.addLast(25);
		dm.delMidNode(ml.head.next.next);
		ml.printList(ml.head);
	}
}




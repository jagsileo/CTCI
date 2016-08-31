import myLinkedList.*;
import java.util.*;
public class RemoveDups {
	private static void removeDupHash(MyNode h) {
		HashMap<Integer, Boolean> hm = new HashMap<Integer, Boolean>();
		MyNode previous = null;
		while(h != null) {
			if(hm.containsKey(h.data)) {
				previous.next = h.next;
			}
			else {
				hm.put(h.data, true);
				previous = h;
			}
			h = h.next;
		}
		
	}
	
	
	public static void main(String[] args) {
		RemoveDups rd = new RemoveDups();
		MyLinkedList ml = new MyLinkedList();
		ml.addLast(5);
		ml.addLast(5);
		ml.addLast(10);
		ml.addLast(15);
		ml.addLast(20);
		ml.addLast(20);
		ml.printList(ml.head);
		rd.removeDupsHash(ml.head);
		ml.printList(ml.head);
	}
}








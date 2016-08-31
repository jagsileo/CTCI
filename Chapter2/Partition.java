import myLinkedList.*;
import java.util.*;
public class Partition {
	private MyNode partitionList(MyNode n, int x) {
		MyNode buf = null;
		MyNode bufHead = null;
		MyNode head = null;
		MyNode curr = null;
		while(n != null) {
			MyNode next = n.next;
			n.next = null;
			if(n.data < x) {
				if(head == null) {
					head = n;
					curr = head;
				} 
				else {
					curr.next = n;
					curr = curr.next;
				}
			}
			else {
				if(buf == null) {
					bufHead = n;					
					buf = bufHead;
				}
				else {
					buf.next = n;
					buf = buf.next;
				}
			}
			n = next;			
		}

		if(head == null) {
			return bufHead;
		}
		curr.next = bufHead;
		return head;
	}

	public static void main(String[] args) {
		MyLinkedList ml = new MyLinkedList();
		Partition p =new Partition();
		ml.addLast(3);
		ml.addLast(5);
		ml.addLast(8);
		ml.addLast(5);
		ml.addLast(10);
		ml.addLast(2);
		ml.addLast(1);		
		ml.printList(ml.head);
		ml.head = p.partitionList(ml.head, 5);
		ml.printList(ml.head);

	}
}














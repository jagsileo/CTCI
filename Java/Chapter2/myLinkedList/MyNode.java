package myLinkedList;
import java.util.*;
public class MyNode {
	public int data;
	public MyNode next;

	public MyNode() {}
	public MyNode(int d) {
		data = d;
		next = null;
	}
	public MyNode(int d, MyNode n) {
		data =d;
		next = n;
	}	

}



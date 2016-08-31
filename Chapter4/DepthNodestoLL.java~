import myTree.*;
//import myLinkedList.*;
import java.util.*;

public class DepthNodestoLL {
	public static void depthToLL(BTNode n, ArrayList<LinkedList<BTNode>> arr, int level) {
		if(n == null) {
			return;
		}

		if(arr.size() == level) {
			LinkedList l = new LinkedList();
			l.add(n);
			arr.add(l);
		}
		else if(arr.size() > level) {
			LinkedList last = arr.get(level);
			last.add(n);
		}
		else if(arr.size() < level) {
			return;
		}

		depthToLL(n.getLeft(), arr, level + 1);
		depthToLL(n.getRight(), arr, level + 1);		

	}




	public static void main(String[] args) {
		BTNode root = new BTNode(5);
		root.setLeft(new BTNode(3));
		root.setRight(new BTNode(7));
		root.getLeft().setLeft(new BTNode(2));
		root.getLeft().setRight(new BTNode(4));
		root.getRight().setLeft(new BTNode(6));
		root.getRight().setRight(new BTNode(8));
		root.printBTNode();
		ArrayList<LinkedList<BTNode>> arr = new ArrayList<LinkedList<BTNode>>();	 
		depthToLL(root, arr, 0);
		for(int i = 0; i<arr.size(); i++) {
			LinkedList ll = arr.get(i);
			for(int j = 0; j<ll.size(); j++) {
				BTNode b = (BTNode)ll.get(j);
				b.printBTNode();
			}
		}
	}
}








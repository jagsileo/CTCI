import myTree.*;
public class FindNext {
	public static MyBTNode findNextInOrder(MyBTNode n) {
		if(n == null) {
			return null;
		}
		if(n.right == null) {
			return null;
		}
		if(n.right.left == null) {
			return n.right;
		}
		else {
			MyBTNode l = n.right.left;
			while(l.left != null) {
				l = l.left;
			}
			return l;			
		}
	}

	public static void main(String[] args) {
		MyBTNode root1 = new MyBTNode(20);
		root1.buildBSTInOrder(new MyBTNode(7));
		root1.buildBSTInOrder(new MyBTNode(4));
		root1.buildBSTInOrder(new MyBTNode(12));
		root1.buildBSTInOrder(new MyBTNode(10));
		root1.buildBSTInOrder(new MyBTNode(9));
		root1.buildBSTInOrder(new MyBTNode(8));
		root1.buildBSTInOrder(new MyBTNode(22));
		root1.buildBSTInOrder(new MyBTNode(21));
		root1.buildBSTInOrder(new MyBTNode(23));	
		
		System.out.println(findNextInOrder(root1.right).data);
	}
}




import myTree.*;
public class CheckBST {
	public static void main(String[] args) {
		MyBTNode root = new MyBTNode(5);
		MyBTNode root1 = new MyBTNode(5);
		root.buildBSTInOrder(new MyBTNode(3));
		root.buildBSTInOrder(new MyBTNode(1));
		root.buildBSTInOrder(new MyBTNode(4));
		root.buildBSTInOrder(new MyBTNode(6));
		root.buildBSTInOrder(new MyBTNode(7));
		root.buildBSTInOrder(new MyBTNode(8));
		
		root1.insert(new MyBTNode(3));
		root1.insert(new MyBTNode(1));
		root1.insert(new MyBTNode(4));
		root1.insert(new MyBTNode(6));
		root1.insert(new MyBTNode(7));
		root1.insert(new MyBTNode(8));
		
		System.out.println(root.isBST());
		System.out.println(root1.isBST());
	}
}


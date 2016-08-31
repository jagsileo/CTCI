import myTree.*;
import java.util.*;
public class BuildBST {
	public static BTNode buildBST(int[] arr, int start, int end) {
		if(end < start) {
			return null;
		}
		int mid = (start + end)/2;
		BTNode root = new BTNode(arr[mid]);		
		root.setLeft(buildBST(arr, start, mid-1));
		root.setRight(buildBST(arr, mid + 1, end));
		return root;		
	}


	public static void main(String[] args) {
		int[] arr = new int[7];
		for(int i = 0; i<7; i++) {
			arr[i] = i+5;
		}

		BTNode root = buildBST(arr, 0 , arr.length-1);
		root.printBTNode();
	}
}












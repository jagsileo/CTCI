import myGraph.*;
import java.util.*;
public class CompileDependency {
	public static int compileOrder(GraphNode p, GraphNode[] order, int count) {
		if(p != null) {
			if(!p.hasNeighbor() && !p.visited) {
				order[count] = p;				
				p.visited = true;
				count++;								
			}

			else if(p.hasNeighbor()){
				for(int i = 0; i<p.getNeighborsCount(); i++) {
					int nCount = compileOrder(p.neighbors[i], order, count);
					order[nCount] = p.neighbors[i];
					count = nCount;
				}
			}		
		}
		return count;
	}

	public static void main(String[] args) {

		GraphNode f = new GraphNode("f");
		GraphNode a = new GraphNode("a");
		GraphNode b = new GraphNode("b");
		GraphNode d = new GraphNode("d");
		GraphNode c = new GraphNode("c");
		GraphNode e = new GraphNode("e");

		a.setNeighbor(f);
		b.setNeighbor(f);
		d.setNeighbor(a);
		d.setNeighbor(b);
		d.setNeighbor(c);

		GraphNode[] ord = new GraphNode[6];
		int fCount = compileOrder(a, ord, 0);
		if(fCount == 6) {

			for(int i = 0; i<ord.length; i++) {
				System.out.println(ord[i].name);
			}
		}
	}
}



















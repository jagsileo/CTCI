import java.util.*;
public class Graph {
	private GraphNode origin;
	private int numNodes;

	public int getNumNodes() {
		return this.numNodes;
	}

	public GraphNode getOrigin() {
		return this.origin;
	}

	public void setOrigin(GraphNode n) {
		this.origin = n;		
	}

	public GraphNode getNodeFromVertex(String s) {
		GraphNode n = this.getOrigin();
		GraphNode result = null;
		if(n.getName() == s) {
			result = n;
		}
		else {
			for(int i = 0; i<n.getNeighbors().size(); i++) {
				if(n.getNeighbors().get(i).getName() == s) {
					result = n.getNeighbors().get(i);
				}
			}
		}
		return result;
	}

	public Map bfs(GraphNode o) {
		Queue<GraphNode> q = new LinkedList<GraphNode>();
		Map<GraphNode, Boolean> visit = new HashMap<GraphNode, Boolean>();
		q.add(o);
		int ctr = 0;
		while(!q.isEmpty()) {
			GraphNode curr = q.remove();

			if(!visit.containsKey(curr)) {
				visit.put(curr, true);
			}
			ctr++;
			if(ctr < o.getNeighbors().size())
				q.add(o.getNeighbors().get(ctr));			
		}
		return visit;

	}
}






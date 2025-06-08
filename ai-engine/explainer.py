import sys
import json

definitions = {
    "hashmap": {
        "definition": "A hashmap is a data structure used to store key-value pairs for fast lookup.",
        "analogy": "Think of it like a phone book — you search for a name (key) to get the number (value)."
    },
    "recursion": {
        "definition": "Recursion is a programming technique where a function calls itself to solve smaller instances of a problem.",
        "analogy": "It's like looking into a mirror that's facing another mirror: each reflection contains a smaller version of the previous one."
    },
    "pointers": {
        "definition": "Pointers are variables that store memory addresses, typically of another variable.",
        "analogy": "A pointer is like a signpost that tells you where to find something, rather than the thing itself."
    }
}

def explain(concept):
    concept = concept.lower()
    if concept in definitions:
        return definitions[concept]
    else:
        return {
            "definition": f"'{concept}' is a computer science concept that involves understanding its core principles and applications in programming and problem-solving.",
            "analogy": f"Think of '{concept}' as a fundamental building block in computing, much like how a foundation is essential for constructing a building."
        }

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print(json.dumps({"error": "No concept provided."}))
        sys.exit(1)
    concept = sys.argv[1]
    result = explain(concept)
    print(json.dumps(result))

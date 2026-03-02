// src/mini/CounterRight.jsx

import { useState } from "react";

// Counter is now "dumb" — it owns NO state
// It receives its value and reports clicks upward via props
function Counter({ label, value, onIncrement }) {
  //                 ↑            ↑
  //         value from parent   function from parent to call when clicked

  return (
    <div
      style={{
        background: "#1a1a1a",
        padding: "1.5rem",
        borderRadius: "8px",
        textAlign: "center",
        width: "150px",
        display: "inline-block",
        margin: "1rem",
      }}
    >
      <p style={{ fontSize: "0.85rem", color: "#888" }}>{label}</p>
      <p style={{ fontSize: "2.5rem", fontWeight: "bold", color: "#fff" }}>
        {value}
      </p>
      <button
        onClick={onIncrement} // just calls what the parent gave it
        style={{
          padding: "0.4rem 1rem",
          background: "#333",
          color: "#fff",
          border: "none",
          borderRadius: "4px",
          cursor: "pointer",
        }}
      >
        +1
      </button>
    </div>
  );
}

// Parent owns ALL the state — it's the single source of truth
export default function CounterRight() {
  // Three separate pieces of state, one per counter
  const [apples, setApples] = useState(0);
  const [bananas, setBananas] = useState(0);
  const [mangoes, setMangoes] = useState(0);

  // Reset All is now trivial — parent controls everything
  function resetAll() {
    setApples(0);
    setBananas(0);
    setMangoes(0);
  }

  return (
    <div>
      <h2>Counter Trio (Right Way)</h2>

      <button
        onClick={resetAll}
        style={{
          padding: "0.6rem 1.5rem",
          background: "#e50914",
          color: "#fff",
          border: "none",
          borderRadius: "6px",
          cursor: "pointer",
          marginBottom: "1rem",
        }}
      >
        Reset All
      </button>

      <div>
        {/*
          Parent passes DOWN:
          - the current value (count)
          - a function to update that specific counter (onIncrement)

          Counter just calls onIncrement when clicked.
          It doesn't know or care what that function does.
        */}
        <Counter
          label="Apples 🍎"
          value={apples}
          onIncrement={() => setApples(apples + 1)}
        />
        <Counter
          label="Bananas 🍌"
          value={bananas}
          onIncrement={() => setBananas(bananas + 1)}
        />
        <Counter
          label="Mangoes 🥭"
          value={mangoes}
          onIncrement={() => setMangoes(mangoes + 1)}
        />
      </div>

      {/* Derived value — calculated from state, not stored as state */}
      <p style={{ marginTop: "1rem", color: "#888" }}>
        Total fruit: {apples + bananas + mangoes}
      </p>
    </div>
  );
}
// ```

// ---

// ## The Diagram That Makes It Click
// ```
//         CounterRight (PARENT)
//         ├── state: apples  = 2
//         ├── state: bananas = 0      ← single source of truth
//         ├── state: mangoes = 5
//         │
//         ├──→ Counter (apples)
//         │     receives: count=2, onIncrement=fn
//         │     when clicked: calls fn → parent updates apples
//         │
//         ├──→ Counter (bananas)
//         │     receives: count=0, onIncrement=fn
//         │
//         └──→ Counter (mangoes)
//               receives: count=5, onIncrement=fn

// Data flows DOWN as props ↓
// Events flow UP as function calls ↑

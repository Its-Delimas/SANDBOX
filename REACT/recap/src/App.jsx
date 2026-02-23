import { Fragment } from "react";

function App() {
  return (
    <Fragment>
      <h2>Hello World</h2>
      <Null />
      <JSX />
    </Fragment>
  );
}

//returning a NULL component
function Null() {
  return <>{/*   null  */}</>;
}
// JSX
function JSX() {
  const myElement = <h1>this is JSX</h1>;

  // embed JavaScript
  const lowercaseClass = "text-lowercase";
  const text = "Hello World";
  const New = <h2 className={lowercaseClass}>{text}</h2>;

  return myElement;
}

export default App;

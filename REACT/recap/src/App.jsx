// Props from Parent Components to Children
function App() {
  return (
    <ChildComponent
      name="Spencer"
      age={19}
      hobbies={["Read Books", "Drink Coffee"]}
      occupation="Software Engineering"
    />
  );
}

function ChildComponent(prop) {
  return (
    <>
      <h1>Hello, my name is {prop.name}</h1>
      <p>
        I am {prop.age} years old...
        <br />
        my hobbies are : {prop.hobbies} <br />
        Occupation : {prop.occupation}
      </p>
    </>
  );
}
export default App;

// inline rendering &&

export default function App() {
  const NewEmail = 2;

  return (
    <>
      <h1>Hello There</h1>

      {NewEmail > 0 && <h2>You have {NewEmail} new emails in your inbox</h2>}
    </>
  );
}

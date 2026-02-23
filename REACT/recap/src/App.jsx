// how to render a list
export default function App() {
  const users = [
    { id: 1, name: "Spencer", role: "Frontend Developer" },
    { id: 2, name: "Delimas", role: "UI & UX Designer" },
    { id: 3, name: "Bangoya", role: "Backend Developer" },
  ];

  return (
    <>
      <ul>
        {users.map((user) => {
          return (
            <li key={user.id}>
              Name : {user.name} ~Role: {user.role}
            </li>
          );
        })}
      </ul>
    </>
  );
}

// Building Forms in React
import { useState } from "react";

// import styles from "styles.module.css";

export default function Form() {
  const [username, setUsername] = useState();
  // ~create a state to store the error messages
  const [usernameError, setUsernameError] = useState();

  // function to handle submit
  const handleSubmit = (e) => {
    e.preventDefault();

    if (usernameError) {
      alert("Unable to Submit: Form Contains errors");
    } else {
      alert(username);
    }
  };
  //perfoming form validation

  const handleUsername = (e) => {
    const { value } = e.target;
    setUsername(value);
    // validate username value
    if (value.length <= 6) {
      setUsernameError("Username length must be more than 6 characters");
    } else {
      setUsernameError();
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      Username:
      <input
        type="text"
        name="username"
        value={username}
        onChange={handleUsername}
      />
      <p>{usernameError}</p>
      <button type="submit">Submit</button>
    </form>
  );
}

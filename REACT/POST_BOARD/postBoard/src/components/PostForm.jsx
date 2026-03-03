import { useState } from "react";
import { createPost } from "../API/postService";

function PostForm({ onPostCreated }) {
  const [title, setTitle] = useState("");
  const [body, setBody] = useState("");
  const [submitting, setSubmitting] = useState(false);
  const [error, seterror] = useState(null);

  async function handleSubmit() {
    if (!title.trim() || !body.trim()) {
      seterror("Both Title and Body are required");
      return;
    }

    setSubmitting(true);
    seterror(null);

    try {
      const newPost = await createPost(title, body);
      onPostCreated(newPost);
      setTitle("");
      setBody("");
    } catch (err) {
      console.error(err);
      seterror("Failed to Create Post");
    } finally {
      setSubmitting(false);
    }
  }

  return (
    <div
      style={{
        background: "#1a1a1a",
        borderRadius: "8px",
        padding: "1.5rem",
        marginBottom: "2rem",
      }}
    >
      <h2 style={{ marginTop: 0 }}>New Post</h2>
      <input
        type="text"
        value={title}
        onChange={(e) => setTitle(e.target.value)}
        style={{
          display: "block",
          width: "100%",
          padding: "0.75rem",
          marginBottom: "0.75rem",
          borderRadius: "6px",
          border: "1px solid #333",
          background: "#0f0f0f",
          color: "#fff",
          fontSize: "1rem",
          boxSizing: "border-box",
        }}
        placeholder="Post Title..."
      />
      <textarea
        value={body}
        onChange={(e) => setBody(e.target.value)}
        placeholder="whats on your mind..."
        rows={4}
        style={{
          display: "block",
          width: "100%",
          padding: "0.75rem",
          marginBottom: "0.75rem",
          borderRadius: "#0f0f0f",
          color: "#fff",
          fontSize: "1rem",
          boxSizing: "border-box",
          resize: "vertical",
        }}
      />
      {error && (
        <p style={{ color: "#e50914", margin: "0 0 0.75rem" }}>⚠ {error}</p>
      )}
      <button
        onClick={handleSubmit}
        disabled={submitting}
        style={{
          padding: "0.75rem 1.5rem",
          borderRadius: "8px",
          border: "none",
          background: submitting ? "#555" : "#238636",
          color: "#fff",
          fontSize: "1rem",
          cursor: submitting ? "not-allowed" : "pointer",
        }}
      >
        {submitting ? "posting..." : "post"}
      </button>
    </div>
  );
}
export default PostForm;

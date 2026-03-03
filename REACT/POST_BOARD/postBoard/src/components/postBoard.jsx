import { useState, useEffect } from "react";
import { getPosts } from "../API/postService";
import PostCard from "./PostCard";
import PostForm from "./PostForm";

export default function PostBoard() {
  const [posts, setPosts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [show, setShow] = useState(true);

  useEffect(() => {
    const fetchPosts = async () => {
      try {
        const data = await getPosts();
        setPosts(data);
      } catch (err) {
        setError("Failed to load posts");
      } finally {
        setLoading(false);
      }
    };
    fetchPosts();
  }, []);

  function handlePostCreate(newPost) {
    setPosts((prevPosts) => [newPost, ...prevPosts]);
  }
  function handleDelete(id) {
    setPosts(posts.filter((post) => post.id !== id));
  }

  return (
    <div style={{ padding: "2rem", maxWidth: "600px", margin: "0 auto" }}>
      <h1>Post Board</h1>
      <PostForm onPostCreated={handlePostCreate} />
      {loading && <p style={{ color: "#888" }}>Loading posts...</p>}
      {error && <p style={{ color: "#e50914" }}>⚠ {error}</p>}

      {!loading && !error && show && (
        <div>
          <h2 style={{ marginBottom: "1rem" }}>Posts ({posts.length})</h2>
          // In the map — parentheses not curly braces // pass the specific
          post's id to handleDelete
          {posts.map((post) => (
            <div key={post.id} style={{ position: "relative" }}>
              <PostCard post={post} />
              <button
                onClick={() => handleDelete(post.id)}
                style={{
                  position: "absolute",
                  top: "1rem",
                  right: "1rem",
                  background: "#e50914",
                  color: "#fff",
                  border: "none",
                  borderRadius: "4px",
                  padding: "0.25rem 0.75rem",
                  cursor: "pointer",
                }}
              >
                Delete
              </button>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

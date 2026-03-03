import { useState, useEffect } from "react";
import { getPosts } from "../API/postService";
import PostCard from "./PostCard";
import PostForm from "./PostForm";

export default function PostBoard() {
  const [posts, setPosts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

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

  return (
    <div style={{ padding: "2rem", maxWidth: "600px", margin: "0 auto" }}>
      <h1>Post Board</h1>
      <PostForm onPostCreated={handlePostCreate} />
      {loading && <p style={{ color: "#888" }}>Loading posts...</p>}
      {error && <p style={{ color: "#e50914" }}>⚠ {error}</p>}

      {!loading && !error && (
        <div>
          <h2 style={{ marginBottom: "1rem" }}>Posts ({posts.length})</h2>
          {posts.map((post) => (
            <PostCard key={post.id} post={post} />
          ))}
        </div>
      )}
    </div>
  );
}

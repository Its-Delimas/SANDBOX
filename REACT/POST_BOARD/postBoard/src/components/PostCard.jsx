function PostCard({ post }) {
  return (
    <div
      style={{
        background: "1a1a1a",
        borderRadius: "8px",
        padding: "1.25rem",
        marginBottom: "0.75rem",
        borderLeft: "3px solid #58a6ff",
      }}
    >
      <span
        style={{
          fontSize: "0.75rem",
          color: "#888",
          background: "#333",
          padding: "0.2rem 0.5rem",
          borderRadius: "4px",
        }}
      >
        #{post.id}
      </span>
      <h3
        style={{
          margin: "0.5rem 0",
          fontSize: "1rem",
          textTransform: "capitalize",
        }}
      >
        {post.Title}
      </h3>
      <p
        style={{
          color: "#aaa",
          fontSize: "0.875rem",
          margin: 0,
          lineHeight: "1.5",
        }}
      >
        {post.body}
      </p>
    </div>
  );
}

export default PostCard;

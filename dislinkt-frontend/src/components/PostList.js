import PostCard from "./PostCard";

const PostList = (posts) => {

    return (
        <div className="container align-content: center display: flex align-items: center mt-5">
            {posts.posts &&
                (posts.posts).map((post, index) => {
                    return (
                        <PostCard key={index} post={post} />
                    );
                })}
        </div>
    );
}

export default PostList;
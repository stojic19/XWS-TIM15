import PostCard from "./PostCard";

const PostList = (props) => {

    return (
        <div className="container align-content: center display: flex align-items: center mt-5">
            {props.posts &&
                (props.posts).map((post, index) => {
                    return (
                        <PostCard key={index} post={post} buttonClickChanger={props.buttonClickChanger}/>
                    );
                })}
            {props.posts.length==0 && <h3 style={{textAlign: 'center'}}>No posts to show.</h3>}
        </div>
    );
}

export default PostList;
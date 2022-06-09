
import '../css/postCard.css'

const JobOfferPost = (props) => {

    return (
        <div className="container">
            <div className="row">
                <div className="[ col-xs-12 col-sm-offset-1 col-sm-7 ]">
                    <div className="[ panel panel-default ] panel-google-plus">

                        <div className="panel-heading">
                            <img className="[ img-circle pull-left ]" src={require("../images/job-avatar.png")} style={{ height: "50px" }} />
                            <h3>{props.offer.position}</h3>
                            <h5><span>
                                {
                                    (props.offer.timeOfCreation).split("T")[0]
                                }</span> </h5>
                        </div>
                        <div className="panel-body">
                            <p>{props.offer.requirements}</p>
                            <p>{props.offer.description}</p>
                        </div>
                        <div className="panel-footer">
                           
                            {/* <ul className="list-group list-group-flush">
                                {
                                    (post.post.comments).map((comment, index) => {
                                        return (
                                            <Comment index={index} comment={comment}></Comment>
                                        );
                                    })
                                }
                            </ul>
                            <br></br>
                            <input className="form-control form-control-sm" type="text" placeholder="Add comment..."
                                onKeyPress={(ev) => {
                                    if (ev.key === "Enter") {
                                        ev.preventDefault();
                                        addPost(ev.target.value);
                                    }
                                }} /> */}
                        </div>
                    </div>
                </div>
            </div>
            <br></br>
            <br></br>
        </div>
      
    );
}

export default JobOfferPost;
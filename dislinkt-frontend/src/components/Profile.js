import '../css/userProfile.css'

const Profile = (props) => {

    return(
        <div className="container">
            <div className="row">
                <div className="col-md-5">
                    <div className="row">
                        <div className="col-12 bg-white p-0 px-3 py-3 mb-5">
                            <div className="d-flex flex-column align-items-center">
                                <h6>
                                {
                                    props.user.isPrivate ? 'Private' : 'Public'
                                }
                                </h6>
                                <img src={require('../images/user-avatar.png')}/>
                                <p className="fw-bold h4 mt-3">{props.user.name}</p>
                                <p className="text-muted">Software developer</p> {/* sortirati listu experience pa podesiti poslednji posao */}
                                <p className="text-muted mb-3">@{props.user.username}</p>
                                <div className="d-flex ">
                                    <div className="btn btn-primary follow me-2">Follow</div>
                                    <div className="btn btn-outline-primary message">Message</div>
                                </div>
                            </div>
                        </div>
                        <div className="col-12 bg-white px-3 pb-2 ">
                            <h6 className="d-flex align-items-center mb-3 fw-bold py-3 justify-content-center"><i
                                    className="text-info me-2">Skills</i></h6>
                                {
                                    props.skills.map((skill, index)=>{
                                        return(
                                            <p style={{textAlign: "center"}} key={index}>{skill}</p>
                                            
                                        )
                                    })
                                }
                        </div>
                        <div className="col-12 bg-white px-3 pb-2 ">
                            <h6 className="d-flex align-items-center mb-3 fw-bold py-3 justify-content-center"><i
                                    className="text-info me-2">Interests</i></h6>
                                {
                                    props.interests.map((interest, index)=>{
                                        return(
                                            <p style={{textAlign: "center"}} key={index}>{interest}</p>
                                            
                                        )
                                    })
                                }
                        </div>
                    </div>
                </div>
                <div className="col-md-7 ps-md-4">
                    <div className="row">
                        <div className="col-12 bg-white p-0 px-2 pb-3 mb-3">
                            <div className="d-flex align-items-center justify-content-between border-bottom">
                                <h4>About</h4>
                            </div>
                            <div className="d-flex align-items-center justify-content-between border-bottom">
                                <p className="py-2">Email</p>
                                <p className="py-2 text-muted">{props.user.email}</p>
                            </div>
                            <div className="d-flex align-items-center justify-content-between border-bottom">
                                <p className="py-2">Mobile</p>
                                <p className="py-2 text-muted">{props.user.telephoneNo}</p>
                            </div>
                            <div className="d-flex align-items-center justify-content-between border-bottom">
                                <p className="py-2">Date of birth</p>
                                <p className="py-2 text-muted">{props.user.dateOfBirth}</p>
                            </div>
                            <div className="d-flex align-items-center justify-content-between border-bottom">
                                <p className="py-2">Gender</p>
                                <p className="py-2 text-muted">{props.user.gender}</p>
                            </div>
                            <div className="d-flex align-items-center justify-content-between border-bottom">
                                <p className="py-2">Biography</p>
                                <p className="py-2 text-muted">{props.user.biography}</p>
                            </div>
                        </div>
                        <br></br>
                        <div className="col-12 bg-white p-0 px-2 pb-3 mb-3">
                            <div className="d-flex align-items-center justify-content-between border-bottom">
                                <h4>Work experience</h4>
                            </div>
                            {
                                (props.experience).map((ex, index)=>{
                                   return(
                                    <div key={index} className="d-flex align-items-center justify-content-between border-bottom">
                                        <p className="py-2">{ex.startDate} - {ex.endDate}</p>
                                        <p className="py-2">{ex.companyName} </p>
                                        <p className="py-2 text-muted">{ex.jobTitle}</p>
                                    </div>
                                   )
                                })
                            }
                        </div>
                        <br></br>
                        <div className="col-12 bg-white px-3 mb-3 pb-3">
                            <div className="d-flex align-items-center justify-content-between border-bottom">
                                <h4>Education</h4>
                            </div>
                            {
                                props.education.map((edu, index)=>{
                                    return(
                                        <div key={index} className="d-flex align-items-center justify-content-between border-bottom">
                                            <p className="py-2">{edu.startDate} - {edu.endDate}</p>
                                            <p className="py-2 text-muted">{edu.institutionName}</p>
                                        </div>
                                    );
                                })
                            }
                        </div>
                        
                    </div>
                </div>
            </div>
        </div>
    );
}

export default Profile;
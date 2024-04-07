import "./Body.css"

function Body() {
    return (
        <div className="body flex items-center justify-center">
            <div className="text-center font-mono">
                <h1 className="uppercase">Welcome Guyzz</h1>
                <br />
                <p>Welcome to the Skill Map project! We're excited to have you onboard as we build a platform for organizations to efficiently manage their employees' skills.</p>
                <p>Your expertise in frontend development is crucial to creating an intuitive user experience. Together, let's revolutionize skill management!</p>
                <p>Happy coding! 🚀</p>
                <br />
                <p>Best regards,<br /> <span className="font-semibold uppercase">SkillMap Project Admins</span></p>
            </div>
        </div>
    )
}

export default Body
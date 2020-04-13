

const AuthItem = () => {
    return (
        <>
            <div style={{
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center'
            }}>
                <img src="/images/user-placeholder-circle-1.png"
                     style={{
                         maxWidth: 25
                     }}
                     alt=""
                />
                <span style={{
                    paddingLeft: '10px'
                }}>
                    Maybe come in?
                </span>
            </div>
        </>
    )
};

export default AuthItem;

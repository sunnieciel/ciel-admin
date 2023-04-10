export const InfoBox = ({img, title, num, clas}) => {
    return <div className={clas}>
        <a href={process.env.IMG_PREFIX + img} target={'_blank'}>
            <img className={'s-icon-64'} src={process.env.IMG_PREFIX + img} alt=""/>
        </a>
        <div className={'info-content ml-12'}>
            <span className={'info-text block mt-6 fs-16'}>{title}</span>
            <span className={'info-num block mt-6 strong'}>{num}</span>
        </div>
    </div>
}
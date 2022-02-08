import {bgGround} from "../utils/utils";

export const TableTd = ({value, type, items, minWidth}) => {
    switch (type) {
        case 'img':
            return <td style={{minWidth: minWidth}}><a href={value} target={'_blank'} rel="noreferrer"><img src={value} alt={`Not found`}/></a></td>
        case "select":
            let info = items.filter(i => i.value === value)[0]
            if (!info) return <td><span>{value}</span></td>
            return <td style={{minWidth: minWidth}}><span
                style={{
                    background: bgGround(value, info.bgGround ? info.bgGround : ''),
                    padding: "0.1em 0.4em",
                    float: "left"
                }}>{info.label}</span></td>
        default:
            return <td style={{minWidth: minWidth}}><span>{value}</span></td>
    }
}
import {bgGround} from "../utils/utils";


export const DetailsLi = ({type, step, label, field, height, items, details, onDetailsChange, placeholder}) => {
    let title = label ? label : field
    switch (type) {
        case 'select_multiple':
        case 'select':
            return <li>
                {title}
                <select multiple={type === 'select_multiple'} value={details[field]} onChange={e => onDetailsChange(e, field)}>
                    <option value={''}>Chose</option>
                    {items.map(item =>
                        <option key={item.value}
                                style={{background: bgGround(item.value, item.bgGround)}}
                                value={item.value}>
                            {item.label}
                        </option>)}
                </select>
            </li>
        case "number":
            return <li>{title}<input type={"number"} step={step} value={details[field] || ''} onChange={e => onDetailsChange(e, field)} placeholder={placeholder}/></li>
        case 'textarea':
            return <li>{title}<textarea rows={height} value={details[field] || ''} onChange={e => onDetailsChange(e, field)} placeholder={placeholder}/></li>
        default:
            return <li>{title}<input value={details[field] || ''} onChange={e => onDetailsChange(e, field)} placeholder={placeholder}/></li>
    }
}

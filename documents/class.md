### 常用clas

登陆后台后在 [http://localhost:1211/to/document](http://localhost:1211/to/document) 进行查看
![](class1.png)
![](class2.png)

## 结构class

### box

一般要写新的一个模块使用`.box` 进行包装

```css
.box {
    border-radius: var(--box-border-radius);
    margin-top: var(--box-spacing);
    box-shadow: var(--elevation-box-shadow-2);
    opacity: 0.99;
}
```

### cell

结合`.box`使用，一般包含在`.box`内。

```css
.cell, .cell-content {
    border-top-left-radius: var(--box-border-radius);
    border-top-right-radius: var(--box-border-radius);
    line-height: 1.5;
    padding: 10px;
    font-size: 14px;
    min-height: 3px;
}
```

## 基础class

```css
.w-80 {
    width: 80px !important;
}

.w-90 {
    width: 90px !important;
}


.text-right {
    text-align: right !important;
}

.m-3 {
    margin: 3px !important;
}

.ml-3 {
    margin-left: 3px !important;
}

.ml-12 {
    margin-left: 12px !important;
}

.mr-3 {
    margin-right: 3px !important;
}

.mr-12 {
    margin-right: 12px !important;
}

.mt-2 {
    margin-top: 2px !important;
}

.mt-3 {
    margin-top: 3px !important;
}

.mt-5 {
    margin-top: 5px !important;
}

.mt-10 {
    margin-top: 10px !important;
}

.mr-auto {
    margin-right: auto !important;
}

.min-w-80 {
    min-width: 80px !important;
}

.p-0 {
    padding: 0 !important;
}

.p-3 {
    padding: 3px !important;
}

.p-10 {
    padding: 10px !important;
}

.pl-20 {
    padding-left: 20px !important;
}

.pb-0 {
    padding-bottom: 0 !important;
}
```

## 字体

```css
.fs-12 {
    font-size: 12px !important;
}

.fs-13 {
    font-size: 13px !important;
}

.fs-14 {
    font-size: 14px !important;
}

.fs-15 {
    font-size: 15px !important;
}

.fs-16 {
    font-size: 16px !important;
}

.strong {
    font-weight: 700 !important;
}

```

## flex

```css
/*flex*/
.flex {
    display: flex;
    align-items: start;
    flex-wrap: wrap;
}

.flex-1 {
    flex: 1;
}

.flex-center {
    display: flex;
    align-items: center;
    flex-wrap: wrap;
}

.flex-baseline {
    display: flex;
    align-items: baseline;
    flex-wrap: wrap;
}

.flex-between {
    display: flex;
    align-items: center;
    justify-content: space-between;
    flex-wrap: wrap;
}

.flex-inline {
    display: inline-flex;
    flex-wrap: wrap;
}
```


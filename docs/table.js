/**
 * table.js
 * 
 * @author caixw <https://github.com/caixw>
 * @license MIT <https://opensource.org/licenses/MIT>
 */


class Table {
  constructor(...cols) {
    this.table = document.createElement('table')

    const thead = document.createElement('thead')
    this.table.appendChild(thead)

    const tr = document.createElement('tr')
    thead.appendChild(tr)

    for(let index in cols) {
      const col = cols[index]

      const th = document.createElement('th')
      if (typeof col == 'object') {
        th.textContent = col[0]
        th.colSpan = col[1]
      }else{
        th.textContent = col
      }

      tr.appendChild(th)
    }

    this.tbody = document.createElement('tbody')
    this.table.appendChild(this.tbody)
  }

  append(...cols) {
    console.log(...cols)
    const tr = document.createElement('tr')
    this.tbody.appendChild(tr)

    for(let index in cols) {
      const td = document.createElement('td')
      td.textContent = cols[index]

      tr.appendChild(td)
    }
  }

  inst() {
    return this.table
  }
}

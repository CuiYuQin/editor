export default (md) => {
  const temp = md.renderer.rules.fence.bind(md.renderer.rules)
  md.renderer.rules.fence = (tokens, idx, options, env, slf) => {
    const token = tokens[idx]
    if (token.info === 'ggb') {
        const code = token.content.trim()
        // return `<div id="ggbApplet" class="ggbApplet"></div><div id="ggbid" class="ggbid">${code}</div>`
        return `<div id="ggbid" class="ggbid">${code}</div>`
    }
    return temp(tokens, idx, options, env, slf)
  }
}
export default (md) => {
  const temp = md.renderer.rules.fence.bind(md.renderer.rules)
  md.renderer.rules.fence = (tokens, idx, options, env, slf) => {
    const token = tokens[idx]
    if (token.info === 'ggb') {
        const code = token.content.trim()
        return `<div id="ggbApplet"></div><div id="ggbid">${code}</div>`
    }
    return temp(tokens, idx, options, env, slf)
  }
}
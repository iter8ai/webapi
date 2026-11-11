import { renderToString } from 'react-dom/server'
import { expect, test } from 'vitest'
import App from './App'

test('App renders its heading', () => {
  expect(renderToString(<App />)).toContain('Weather')
})

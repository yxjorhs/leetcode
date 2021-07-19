abstract class ATest {
  a (): number {
    return 1
  }
}

interface IATest {
  a():number
}

class Test implements IATest {
  a() {
    return 1
  }
}

function run() {
  const t = new Test()
  console.log('------t.a()--------', t.a())
}

run()
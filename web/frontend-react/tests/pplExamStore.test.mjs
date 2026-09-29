import test from 'node:test';
import assert from 'node:assert/strict';
import { readFileSync } from 'node:fs';
import { createRequire } from 'node:module';
import vm from 'node:vm';
import ts from 'typescript';
const require=createRequire(import.meta.url);
const compiled=ts.transpileModule(readFileSync(new URL('../src/store/pplExamStore.ts',import.meta.url),'utf8'),{compilerOptions:{module:ts.ModuleKind.CommonJS}}).outputText;
function setup({saved=null,failAt=-1}={}) {
  let time=1000, storage=saved, calls=0;
  const questions=Array.from({length:12},(_,i)=>({id:`q${i}`,options:['faux','juste','autre','dernier'],options_en:['wrong','right','other','last']}));
  const context={exports:{},Date:{now:()=>time},sessionStorage:{getItem:()=>storage,setItem:(_,value)=>{storage=value;}},require:(name)=>{
    if(name==='zustand')return require('zustand');
    if(name==='@/api/client')return {apiClient:{get:async()=>({data:{paper:{id:'human_performance',count:12,minutes:18},questions}})}};
    if(name==='@/api/questions')return {questionsApi:{answer:async(id,letter)=>{
      calls++;if(calls===failAt)throw Error('offline');
      return {data:{correct:letter==='B',correct_answer:'juste',explanation_fr:'Correction',explanation_en:'Explanation'}};
    }}};
    throw Error(name);
  }};
  vm.runInNewContext(compiled,context);
  return {store:context.exports.usePPLExamStore,calls:()=>calls,save:()=>storage,advance:(ms)=>{time+=ms;}};
}
test('no correction before submission; 9/12 passes and blanks receive zero',async()=>{
  const x=setup();await x.store.getState().start('human_performance');
  for(let i=0;i<9;i++)x.store.getState().choose(`q${i}`,'B');
  assert.equal(x.calls(),0);assert.equal(Object.keys(x.store.getState().results).length,0);
  await x.store.getState().finish();
  assert.equal(x.store.getState().completed.human_performance.correct,9);
  assert.equal(x.store.getState().completed.human_performance.passed,true);
  assert.equal(x.store.getState().results.q11.correct,false);
  x.store.getState().choose('q11','B');assert.equal(x.store.getState().answers.q11,undefined);
});
test('8/12 fails and another paper result cannot substitute for it',async()=>{
  const x=setup();await x.store.getState().start('human_performance');
  for(let i=0;i<8;i++)x.store.getState().choose(`q${i}`,'B');
  await x.store.getState().finish();assert.equal(x.store.getState().completed.human_performance.passed,false);
  assert.equal(Object.keys(x.store.getState().completed).length,1);
});
test('deadline survives reload and late choices are rejected',async()=>{
  const x=setup();await x.store.getState().start('human_performance');x.store.getState().choose('q0','A');
  const deadline=x.store.getState().deadline;
  const y=setup({saved:x.save()});assert.equal(y.store.getState().deadline,deadline);
  y.advance(18*60_000);y.store.getState().choose('q0','B');
  assert.equal(y.store.getState().answers.q0,'A');await y.store.getState().finish();
  assert.equal(y.store.getState().phase,'finished');
});
test('failed grading freezes choices and retry does not resubmit completed corrections',async()=>{
  const x=setup({failAt:3});await x.store.getState().start('human_performance');
  x.store.getState().choose('q0','B');await x.store.getState().finish();
  assert.equal(x.store.getState().phase,'grading');assert.equal(Object.keys(x.store.getState().results).length,2);
  x.store.getState().choose('q2','B');assert.equal(x.store.getState().answers.q2,undefined);
  const y=setup({saved:x.save()});await y.store.getState().finish();
  assert.equal(y.calls(),10);assert.equal(y.store.getState().phase,'finished');
});

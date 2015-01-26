# [buildgradle.github.io](https://buildgradle.github.io)

## AKA: check out my build.gradle!

**Call for build.gradles.**  
You are more than welcome to fork and pull request with your own projects' build.gradle files.

### FAQ

#### Funny project. Why did you create it?

I was starting a new project and figured out, I ask my fellow Android developers, what libraries do they use. A few weeks before that, another Android dev asked me similar question. So it's probable a few (hundred) other may ask their friends the same questions.

The idea is to share `build.gradle` (and other `dot.gradle`) files from your project, so we can see not only what libraries other use, but to learn some tricks, that we would not even think of...

> Remember: we are all junior Groovy devs.

#### Can I share build.gradle from my work project?

You should probably ask your boss first and obfuscate at least applicationId. We won't ask you to sign CLA. If you want to share it secretly, consider putting it in some random directory like `company1/example1` (or `company666` if first is already taken).

#### My boss said I can share these files without modification. How cool is that?

That's pretty awesome. And is very likely beneficial for your company. Think of it this way: you have RxJava + Groovy in your dependencies and some other fancy stuff. What developer would not work for a company that experiments with these sort of things? You can use `description` in [`project.json`](files/mg6maciej/android-maps-extensions/project.json) to brag about it even more.

#### What is this project.json for?

This is some meta data about your project. Currently these fields are read:

* name
* description
* url

You may use them for anything you want. Link to job offer? No problem.


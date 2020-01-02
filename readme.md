#### Description

This repository is an experiment with go plugins. The intended action was to at runtime register commands from an external code.

- The main cli code that load addons is located at `cli`
- The external code is located at `addon-gerard`. 
- Addons should be compiled in `addons` directory

#### How to run

`make run`

#### Problems

- The plugin compiler version must exactly match the program's compiler version. If the program was compiled with 1.11.4, it won't work to compile the plugin with 1.11.5. When distributing a program binary, you must communicate what the compiler version you used is.
- Any packages outside of the standard library that are used by both the plugin and the program must have their versions match exactly. In my case I am using cobra in version v0.0.5, if the plugin version is v0.0.4 it won't work.

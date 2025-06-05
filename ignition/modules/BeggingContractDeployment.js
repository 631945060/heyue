const { buildModule } = require('@nomicfoundation/ignition-core'); 
 
// 构建一个新的 Ignition 模块 
const BeggingContractDeployment = buildModule('BeggingContractDeployment', (m) => { 
    // 定义合约部署 
    const BeggingContract = m.contract('BeggingContract');  
 
    return { 
        BeggingContract 
    }; 
}); 
 
module.exports  = { 
    BeggingContractDeployment 
}; 